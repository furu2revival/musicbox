import "./style.module.css";
import { useMemo } from "react";
import sound from "~/assets/hole.mp3";
import type { Note, Pitch } from "~/model/note";
import style from "./style.module.css";

const PITCHES: Pitch[] = ["C3", "D3", "E3", "F3", "G3", "A4", "B4", "C4"];

type Props = {
	notes: Note[];
	onChange?: (timingIndex: number, pitch: Pitch) => void;
	className?: string;
};
export const SheetTable = ({ notes, onChange, className }: Props) => {
	const audio = useMemo(() => new Audio(sound), []);
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
									return (
										// biome-ignore lint/a11y/useKeyWithClickEvents: <explanation>
										<td
											key={`${pitch}:${
												//ここでは配列の長さは必ず同じでかつ、同一のiは同じNoteを指しているので、iを使っても問題ない
												// biome-ignore lint/suspicious/noArrayIndexKey: <explanation>
												i
											}`}
											onClick={() => {
												audio.play();
												onChange?.(i, pitch);
											}}
										>
											<span className={style["shift-right"]}>
												{note.pitch.filter((p) => p === pitch).length
													? "●"
													: ""}
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
