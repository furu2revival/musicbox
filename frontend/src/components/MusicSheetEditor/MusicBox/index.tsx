import { useMemo } from "react";
import playIcon from "~/assets/play.png";
import resetSound from "~/assets/reset.mp3";
import share from "~/assets/share.png";
import stopIcon from "~/assets/stop.png";
import trashCan from "~/assets/trash_can.png";
import style from "./style.module.css";

type Props = {
	className?: string;
	onShare: () => void;
	onReset: () => void;
	isCharge?: boolean;
	shareDisabled: boolean;
	energy: number;
	onStart: () => void;
	onStop: () => void;
	onLoad: () => void;
};

export const MusicBox = ({
	className,
	onShare,
	onReset,
	isCharge,
	shareDisabled,
	energy,
	onStart,
	onStop,
	onLoad,
}: Props) => {
	const resetAudio = useMemo(() => new Audio(resetSound), []);
	const shakeIntensity = energy * 0.1;
	return (
		<div className={`${style.root} ${className}`}>
			<button
				className={`${style.iconButton} ${shareDisabled ? style.disabled : ""}`}
				type="button"
				onClick={onShare}
				disabled={shareDisabled}
			>
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
				className={style.iconButton}
				type="button"
				onClick={() => {
					if (window.confirm("譜面をリセットしますか？")) {
						resetAudio.play();
						onReset();
					}
				}}
			>
				<img width={32} src={trashCan} alt="" />
			</button>

			<button
				className={style.iconButton}
				onClick={() => onStart()}
				type="button"
			>
				<img src={playIcon} alt="再生" width={32} />
			</button>
			<button
				onClick={() => onStop()}
				className={style.iconButton}
				type="button"
			>
				<img src={stopIcon} alt="ストップ" height={32} />
			</button>
			<button onClick={() => onLoad()} type="button">
				load
			</button>
			{energy}
		</div>
	);
};
