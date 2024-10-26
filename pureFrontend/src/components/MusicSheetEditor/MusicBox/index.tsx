import { Propeller } from "./Propeller";
import style from "./style.module.css";

type Props = {
	className?: string;
	height: number;
};

export const MusicBox = ({ className, height }: Props) => {
	return (
		<div
			className={`${style.root} ${className}`}
			style={{
				height,
			}}
		>
			<button type="button">共有</button>
			<Propeller />
			<button type="button">削除</button>
		</div>
	);
};
