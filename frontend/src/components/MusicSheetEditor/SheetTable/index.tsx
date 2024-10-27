import "./style.module.css";
import holeSound from "~/assets/hole.mp3";
import type { Note, Pitch } from "~/model/note";
import { Hole } from "./Hole";
import style from "./style.module.css";
import { usePeriodicAudioPlayer } from "~/hooks/periodicAudioPlayer";

const PITCHES: Pitch[] = ["C3", "D3", "E3", "F3", "G3", "A4", "B4", "C4"];

type Props = {
	notes: Note[];
	onChange?: (timingIndex: number, pitch: Pitch) => void;
	className?: string;
};
export const SheetTable = ({ notes, onChange, className }: Props) => {
  const { player } = usePeriodicAudioPlayer(holeSound, 0.025, 0.3);

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
								{i + 1}
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
												if (!isHoled) player?.play(0, 1);
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
