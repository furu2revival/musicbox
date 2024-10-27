import { useEffect, useState } from "react";

import type {
	MusicSheetPlayer,
	MusicSheetPlayerInit,
} from "~/features/musicSheetPlayer/musicSheetPlayer";
import { createMusicSheetPlayer } from "~/features/musicSheetPlayer/musicSheetPlayer";
import type { PeriodicAudioPlayer } from "~/features/musicSheetPlayer/periodicAudioPlayer";
import { createPeriodicAudioPlayer } from "~/features/musicSheetPlayer/periodicAudioPlayer";

import windSprintSound from "~/assets/wind_spring.wav";

export const useMusicSheetPlayer = (init: MusicSheetPlayerInit) => {
	const [player, setPlayer] = useState<MusicSheetPlayer | null>(null);
	const [windPlayer, setWindPlayer] = useState<PeriodicAudioPlayer | null>(
		null
	);
	const [oldEnergy, setOldEnergy] = useState<number>(0);

	useEffect(() => {
		(async () => {
			const player = await createMusicSheetPlayer(init);
			setPlayer(player);
		})();
	}, [...Object.values(init)]);

	useEffect(() => {
		(async () => {
			const player = await createPeriodicAudioPlayer(windSprintSound, 0.001);
			setWindPlayer(player);
		})();
	});

	useEffect(() => {
		const listener = () => {
			const newEnergy = player?.energy ?? 0;
			const maxEnergy = player?.maxEnergy;
			if (typeof maxEnergy === "number" && oldEnergy < newEnergy) {
				const volume = oldEnergy / maxEnergy;

				if (volume < 0.5) {
					windPlayer?.play(50, 10);
				} else if (volume < 0.75) {
					windPlayer?.play(100, 10);
				} else if (volume < 0.9) {
					windPlayer?.play(200, 5);
				} else {
					windPlayer?.play(500, 3);
				}
			}

			setOldEnergy(newEnergy);
		};
		player?.addEventListener("energychange", listener);

		return () => {
			player?.removeEventListener("energychange", listener);
		};
	}, [player, windPlayer]);

	const play = () => player?.play();
	const stop = () => player?.stop();
	const setPlayerEnergy = (energy: number) => player?.setEnergy(energy);

	return {
		ready: player !== null && windPlayer !== null,
		play,
		stop,
		energy: player?.energy,
		setEnergy: setPlayerEnergy,
		maxEnergy: player?.maxEnergy,
	};
};
