import { useEffect, useState } from "react";

import {
	createMusicSheetPlayer,
	MusicSheetPlayerInit,
	MusicSheetPlayer,
} from "./musicSheetPlayer";

export const useMusicSheetPlayer = function (init: MusicSheetPlayerInit) {
	const [player, setPlayer] = useState<MusicSheetPlayer | null>(null);
	const [energy, setEnergy] = useState<number>(0);

	useEffect(() => {
		(async () => {
			const player = await createMusicSheetPlayer(init);
			setPlayer(player);
		})();
	}, [...Object.values(init)]);

	useEffect(() => {
		const listener = () => {
			setEnergy(player?.energy || 0);
		};
		player?.addEventListener("energychange", listener);

		return () => {
			player?.removeEventListener("energychange", listener);
		};
	}, [player]);

	const play = () => player?.play();
	const stop = () => player?.stop();
	const setPlayerEnergy = (energy: number) => player?.setEnergy(energy);

	return {
		ready: player !== null,
		play,
		stop,
		energy,
		setEnergy: setPlayerEnergy,
	};
};
