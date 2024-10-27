import { useEffect, useState } from "react";
import {
	type PeriodicAudioPlayer,
	createPeriodicAudioPlayer,
} from "~/features/musicSheetPlayer/periodicAudioPlayer";

export const usePeriodicAudioPlayer = function (
	src: string,
	startTime = 0,
	volume = 0
) {
	const [player, setPlayer] = useState<PeriodicAudioPlayer | null>(null);

	useEffect(() => {
		(async () => {
			const player = await createPeriodicAudioPlayer(src, startTime, volume);
			setPlayer(player);
		})();
	}, []);

	return { player };
};
