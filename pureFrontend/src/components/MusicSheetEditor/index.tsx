import { useState } from "react";
import type { Note } from "../../../../frontend/app/model/note";
import { MusicBox } from "./MusicBox";
import { SheetTable } from "./SheetTable";
import style from "./style.module.css";

export const MusicSheetEditor = () => {
	const [notes, setNotes] = useState<Note[]>([
		// one beat
		{ pitch: ["C3"] },
		{ pitch: ["D3", "F3", "A4"] },
		{ pitch: ["E3", "G3", "B4"] },
		{ pitch: ["F3", "A4", "C4"] },
		// one beat
		{ pitch: ["G3", "B4", "D3"] },
		{ pitch: ["A4", "C4", "E3"] },
		{ pitch: ["B4", "D3", "F3"] },
		{ pitch: ["C4", "E3", "G3"] },
		// one beat
		{ pitch: ["D3", "F3", "A4"] },
		{ pitch: ["E3", "G3", "B4"] },
		{ pitch: ["F3", "A4", "C4"] },
		{ pitch: ["G3", "B4", "D3"] },
		// one beat
		{ pitch: ["A4", "C4", "E3"] },
		{ pitch: ["B4", "D3", "F3"] },
		{ pitch: ["C4", "E3", "G3"] },
		{ pitch: ["D3", "F3", "A4"] },
		// one beat
		{ pitch: ["C3", "E3", "G3"] },
		{ pitch: ["D3", "F3", "A4"] },
		{ pitch: ["E3", "G3", "B4"] },
		{ pitch: ["F3", "A4", "C4"] },
		// one beat
		{ pitch: ["G3", "B4", "D3"] },
		{ pitch: ["A4", "C4", "E3"] },
		{ pitch: ["B4", "D3", "F3"] },
		{ pitch: ["C4", "E3", "G3"] },
		// one beat
		{ pitch: ["D3", "F3", "A4"] },
		{ pitch: ["E3", "G3", "B4"] },
		{ pitch: ["F3", "A4", "C4"] },
		{ pitch: ["G3", "B4", "D3"] },
		// one beat
		{ pitch: ["A4", "C4", "E3"] },
		{ pitch: ["B4", "D3", "F3"] },
		{ pitch: ["C4", "E3", "G3"] },
		{ pitch: ["D3", "F3", "A4"] },
		// one beat
		{ pitch: ["C3", "E3", "G3"] },
		{ pitch: ["D3", "F3", "A4"] },
		{ pitch: ["E3", "G3", "B4"] },
		{ pitch: ["F3", "A4", "C4"] },
		// one beat
		{ pitch: ["G3", "B4", "D3"] },
		{ pitch: ["A4", "C4", "E3"] },
		{ pitch: ["B4", "D3", "F3"] },
		{ pitch: ["C4", "E3", "G3"] },
		// one beat
		{ pitch: ["D3", "F3", "A4"] },
		{ pitch: ["E3", "G3", "B4"] },
		{ pitch: ["F3", "A4", "C4"] },
		{ pitch: ["G3", "B4", "D3"] },
		// one beat
		{ pitch: ["A4", "C4", "E3"] },
		{ pitch: ["B4", "D3", "F3"] },
		{ pitch: ["C4", "E3", "G3"] },
		{ pitch: ["D3", "F3", "A4"] },
		// one beat
		{ pitch: ["C3", "E3", "G3"] },
		{ pitch: ["D3", "F3", "A4"] },
		{ pitch: ["E3", "G3", "B4"] },
		{ pitch: ["F3", "A4", "C4"] },
		// one beat
		{ pitch: ["G3", "B4", "D3"] },
		{ pitch: ["A4", "C4", "E3"] },
		{ pitch: ["B4", "D3", "F3"] },
		{ pitch: ["C4", "E3", "G3"] },
		// one beat
		{ pitch: ["D3", "F3", "A4"] },
		{ pitch: ["E3", "G3", "B4"] },
		{ pitch: ["F3", "A4", "C4"] },
		{ pitch: ["G3", "B4", "D3"] },
		// one beat
		{ pitch: ["A4", "C4", "E3"] },
		{ pitch: ["B4", "D3", "F3"] },
		{ pitch: ["C4", "E3", "G3"] },
		{ pitch: ["D3", "F3", "A4"] },
	]);

	return (
		<div className={style.root}>
			<MusicBox className={style["music-box-position"]} height={100} />
			<SheetTable
				className={style["sheet-table-position"]}
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
