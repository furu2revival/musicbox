import { useState } from "react";
import type { MusicSheet } from "~/model/musicSheet";
import type { Note } from "~/model/note";
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
	const notesInit = Array(maxNotes).fill({ pitch: [] });
	const [notes, setNotes] = useState<Note[]>(musicSheet?.notes ?? notesInit);

	return (
		<div className={`${style.root} ${className ?? ""}`}>
			<MusicBox
				energy={energy}
				onReset={() => setNotes(notesInit)}
				onShare={async () => {
					const postNotes = () => {
						console.log(notes, "posted!");
						return "id_id_id_id_id_id";
					};
					const noteId = postNotes();
					const url = `https://example.com/notes/${noteId}`;
					if (navigator.share)
						await navigator.share({
							title: "テスト",
							text: "ほげ",
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
