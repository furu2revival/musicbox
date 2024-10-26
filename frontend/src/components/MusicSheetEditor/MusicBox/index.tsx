import { Propeller } from "./Propeller";
import style from "./style.module.css";

type Props = {
	className?: string;
	onShare: () => void;
	onReset: () => void;
};

export const MusicBox = ({ className, onShare, onReset }: Props) => {
	return (
		<div className={`${style.root} ${className}`}>
			<button type="button" onClick={onShare}>
				共有
			</button>
			<Propeller />
			<button type="button" onClick={onReset}>
				削除
			</button>
		</div>
	);
};
