import { useMemo } from "react";
import resetSound from "~/assets/reset.mp3";
import share from "~/assets/share.png";
import trashCan from "~/assets/trash_can.png";
import style from "./style.module.css";

type Props = {
	className?: string;
	onShare: () => void;
	onReset: () => void;
	isCharge?: boolean;
	energy: number;
};

export const MusicBox = ({
	className,
	onShare,
	onReset,
	isCharge,
	energy,
}: Props) => {
	const resetAudio = useMemo(() => new Audio(resetSound), []);
	const shakeIntensity = energy * 0.1;
	return (
		<div className={`${style.root} ${className}`}>
			<button className={style.shareButton} type="button" onClick={onShare}>
				<img width={32} src={share} alt="" />
			</button>
			<div
				style={{
					color: "white",
					animation: "shake 0.5s infinite",
					transform: `translate(${shakeIntensity}px, ${shakeIntensity}px)`,
				}}
				className={`${isCharge ? style.rotate : ""}`}
			>
				ぷ
			</div>
			<button
				className={style.deleteButton}
				type="button"
				onClick={() => {
					resetAudio.play();
					onReset();
				}}
			>
				<img width={32} src={trashCan} alt="" />
			</button>
		</div>
	);
};
