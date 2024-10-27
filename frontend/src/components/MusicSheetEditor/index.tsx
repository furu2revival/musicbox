import { useState } from "react";
import { usePostMusicSheet } from "~/hooks/usePostMusicSheet";
import type { MusicSheet } from "~/model/musicSheet";
import { type Note, NoteToResponse } from "~/model/note";
import { MusicBox } from "./MusicBox";
import { SheetTable } from "./SheetTable";
import style from "./style.module.css";
import { useMusicBox } from "~/hooks/furu2Musicbox";

type Props = {
	className?: string;
	maxNotes: number;
	musicSheet: MusicSheet | undefined;
	maxEnergy: number;
};

export const MusicSheetEditor = ({
	className,
	maxNotes,
	musicSheet,
	maxEnergy,
}: Props) => {
	const notesInit: Note[] = Array(maxNotes).fill({ pitch: [] } as Note);
	const [notes, setNotes] = useState<Note[]>(musicSheet?.notes ?? notesInit);
	const { postMusicSheet } = usePostMusicSheet();
	const musicBox = useMusicBox({
		playerInit: {
			musicSheet: {
				id: musicSheet?.id ?? "",
				title: musicSheet?.title ?? "",
				notes,
			},
			beatsPerMinute: 120,
			maxEnergy: maxEnergy ?? 100,
		},
		shakeDetectorInit: {
			shakeDetectInterval: 100,
			accelerationThreshold: 5,
			moveAmountThreshold: 0.2,
		},
	});

	return (
		<>
			<div className={`${style.root} ${className ?? ""}`}>
				<MusicBox
					energy={musicBox.energy}
					onReset={async () => {
						setNotes(notesInit);
					}}
					onShare={async () => {
						const noteId = (
							await postMusicSheet({
								notes: notes.map(NoteToResponse),
								title: "テスト:将来的にこの値は変更されます",
							})
						).musicSheetId;
						const url = `${new URL(window.location.href).origin}?musicSheetId=${noteId}`;

						if (navigator.share)
							await navigator.share({
								title: "ふるふるオルゴール",
								text: "オルゴールオリジナル楽曲",
								url,
							});
						else
							navigator.clipboard
								.writeText(url)
								.then(() =>
									alert(`URLがクリップボードにコピーされました: ${url}`)
								)
								.catch((err) =>
									console.error("クリップボードへのコピーに失敗しました: ", err)
								);
					}}
					isCharge={true}
          shareDisabled={notes.every((note) => note.pitch.length === 0)}
				/>
				<SheetTable
					notes={notes}
					onChange={(timingIndex, pitch) => {
						setNotes((prev) => {
							const newNotes = [...prev];
							newNotes[timingIndex] = {
								pitch: [...newNotes[timingIndex].pitch, pitch],
							};
							return newNotes;
						});
					}}
				/>
			</div>
			<div>
				<div>ready: {musicBox.ready.toString()}</div>
				<div>
					energy: {musicBox.energy} / {musicBox.maxEnergy}
				</div>
				<div>
					<button onClick={() => musicBox.load()}>load</button>
					<button onClick={() => musicBox.play()}>play</button>
					<button onClick={() => musicBox.stop()}>stop</button>
				</div>
			</div>
		</>
	);
};
