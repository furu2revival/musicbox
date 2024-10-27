import { useState } from "react";
import { usePostMusicSheet } from "~/hooks/usePostMusicSheet";
import type { MusicSheet } from "~/model/musicSheet";
import { type Note, NoteToResponse } from "~/model/note";
import { MusicBox } from "./MusicBox";
import { SheetTable } from "./SheetTable";
import style from "./style.module.css";

type Props = {
	className?: string;
	maxNotes: number;
	musicSheet: MusicSheet | undefined;
	energy: number;
};

export const MusicSheetEditor = ({
	className,
	maxNotes,
	musicSheet,
	energy,
}: Props) => {
	const notesInit: Note[] = Array(maxNotes).fill({ pitch: [] } as Note);
	const [notes, setNotes] = useState<Note[]>(musicSheet?.notes ?? notesInit);
	const { postMusicSheet } = usePostMusicSheet();

	return (
		<div className={`${style.root} ${className ?? ""}`}>
			<MusicBox
				energy={energy}
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
								alert(`URLがクリップボードにコピーされました: ${url}`),
							)
							.catch((err) =>
								console.error("クリップボードへのコピーに失敗しました: ", err),
							);
				}}
				isCharge={true}
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
	);
};
