import "./style.module.css";
import { useMemo } from "react";
import holeSound from "~/assets/hole.mp3";
import type { Note, Pitch } from "~/model/note";
import { Hole } from "./Hole";
import style from "./style.module.css";

const PITCHES: Pitch[] = ["C4", "B4", "A4", "G3", "F3", "E3", "D3", "C3"];

type Props = {
	notes: Note[];
	onChange?: (timingIndex: number, pitch: Pitch) => void;
	className?: string;
};
export const SheetTable = ({ notes, onChange, className }: Props) => {
	const audio = useMemo(() => new Audio(holeSound), []);
	return (
		<div className={className}>
			<table
				style={{
					height: "full",
					width: "full",
				}}
			>
				<thead>
					<tr>
						<th />
						{notes.map((_, i) => (
							<th
								key={
									//ここでは配列の長さは必ず同じでかつ、同一のiは同じNoteを指しているので、iを使っても問題ない
									// biome-ignore lint/suspicious/noArrayIndexKey: <explanation>
									i
								}
							>
								{i}
							</th>
						))}
					</tr>
				</thead>
				<tbody>
					{PITCHES.map((pitch) => {
						return (
							<tr key={pitch}>
								<td width={200}>{pitch}</td>
								{notes.map((note, i) => {
									const isHoled = note.pitch.filter((p) => p === pitch).length;
									return (
										// biome-ignore lint/a11y/useKeyWithClickEvents: <explanation>
										<td
											key={`${pitch}:${
												//ここでは配列の長さは必ず同じでかつ、同一のiは同じNoteを指しているので、iを使っても問題ない
												// biome-ignore lint/suspicious/noArrayIndexKey: <explanation>
												i
											}`}
											onClick={() => {
												if (!isHoled) audio.play();
												onChange?.(i, pitch);
											}}
										>
											<span className={style["shift-right"]}>
												{isHoled ? <Hole /> : ""}
											</span>
										</td>
									);
								})}
							</tr>
						);
					})}
				</tbody>
			</table>
		</div>
	);
};
