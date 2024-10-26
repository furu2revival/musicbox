import "./style.module.css";
import type { Note, Pitch } from "~/model/note";
import style from "./style.module.css";

const PITCHES: Pitch[] = ["C3", "D3", "E3", "F3", "G3", "A4", "B4", "C4"];

type Props = {
	notes: Note[];
	onChange?: (timingIndex: number, pitch: Pitch) => void;
	className?: string;
};
export const SheetTable = ({ notes, onChange, className }: Props) => {
	return (
		<div className={className}>
			<table>
				<thead>
					<tr>
						{notes.map(({ pitch }, i) => (
							<th key={pitch.join(",")}>{i}</th>
						))}
					</tr>
				</thead>
				<tbody>
					{PITCHES.map((pitch) => {
						return (
							<tr key={pitch}>
								{notes.map((note, i) => {
									return (
										// biome-ignore lint/a11y/useKeyWithClickEvents: <explanation>
										<td
											// biome-ignore lint/suspicious/noArrayIndexKey: <explanation>
											key={pitch + i}
											onClick={() => {
												console.log("click");
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
