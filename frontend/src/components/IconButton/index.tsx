import type React from "react";

type Props = {
	Icon: React.ReactNode;
	onClick: () => void;
};
export const IconButton = ({ Icon, onClick }: Props) => {
	return (
		<button type="button" onClick={onClick}>
			{Icon}
		</button>
	);
};
