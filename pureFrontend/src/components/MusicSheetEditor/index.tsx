import { useState } from "react";
import type { MusicSheet } from "~/model/musicSheet";
import type { Note } from "../../../../frontend/app/model/note";
import { MusicBox } from "./MusicBox";
import { SheetTable } from "./SheetTable";
import style from "./style.module.css";

type Props = {
	className?: string;
	musicSheet: MusicSheet;
};

export const MusicSheetEditor = ({ className, musicSheet }: Props) => {
	const [notes, setNotes] = useState<Note[]>(musicSheet.notes);

	return (
		<div className={`${style.root} ${className ?? ""}`}>
			<MusicBox
				onReset={() => {
					setNotes(Array(64).fill({ pitch: [] }));
				}}
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
							.then(() => {
								// biome-ignore lint/style/useTemplate: <explanation>
								alert("URLがクリップボードにコピーされました: " + url);
							})
							.catch((err) => {
								console.error("クリップボードへのコピーに失敗しました: ", err);
							});
				}}
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
