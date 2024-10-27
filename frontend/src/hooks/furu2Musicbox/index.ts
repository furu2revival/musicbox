import { useEffect, useState } from "react";

import {
	MusicSheetPlayer,
	MusicSheetPlayerInit,
} from "~/features/musicSheetPlayer/musicSheetPlayer";
import { createMusicSheetPlayer } from "~/features/musicSheetPlayer/musicSheetPlayer";
import type { PeriodicAudioPlayer } from "~/features/musicSheetPlayer/periodicAudioPlayer";
import { createPeriodicAudioPlayer } from "~/features/musicSheetPlayer/periodicAudioPlayer";
import type { ShakeDetectorInit } from "~/features/shakeDetecters";
import { createShakeDetector, ShakeDetector } from "~/features/shakeDetecters";

import windSprintSound from "~/assets/wind_spring.wav";

type Furu2MusicPlayerInit = {
  playerInit: MusicSheetPlayerInit,
  shakeDetectorInit: ShakeDetectorInit,
}

export const useMusicBox = function(init: Furu2MusicPlayerInit) {
  const [player, setPlayer] = useState<MusicSheetPlayer | null>(null);
	const [windPlayer, setWindPlayer] = useState<PeriodicAudioPlayer | null>(null);
  const [shakeDetector, setShakeDetector] = useState<ShakeDetector | null>(null);
  const [energy, setEnergy] = useState<number>(0);
  const [maxEnergy, setMaxEnergy] = useState<number>(0);

  const loadPlayer = async function() {
    const player = await createMusicSheetPlayer(init.playerInit);
    setPlayer(player);
  }
  const loadWindPlayer = async function() {
    const player = await createPeriodicAudioPlayer(windSprintSound, 0.001);
    setWindPlayer(player);
  }
  const loadShakeDetector = async function() {
    const detector = await createShakeDetector(init.shakeDetectorInit);
    setShakeDetector(detector);
  }
  const load = function() {
    loadPlayer();
    loadWindPlayer();
    loadShakeDetector();
  }

  useEffect(() => {
    const handleShake = () => {
      if(!player) return;

      const energy = player.energy;
			const maxEnergy = player.maxEnergy;

      const volume = energy / maxEnergy;

      if (volume < 0.5) {
        windPlayer?.play(50, 10);
      } else if (volume < 0.75) {
        windPlayer?.play(100, 10);
      } else if (volume < 0.9) {
        windPlayer?.play(200, 5);
      } else {
        windPlayer?.play(500, 3);
      }

      player?.setEnergy(energy + 1);
    };

    shakeDetector?.addEventListener('shake', handleShake);

    return () => {
      shakeDetector?.removeEventListener('shake', handleShake);
    }
  }, [player, shakeDetector]);
  
  useEffect(() => {
    if(!player) return;

    const handleEnergyChange = () => {
      setEnergy(player.energy);
      setMaxEnergy(player.maxEnergy);
    };
    player.addEventListener('energychange', handleEnergyChange);
    return () => {
      player.removeEventListener('energychange', handleEnergyChange);
    }
  }, [player]);

  return {
    ready: player !== null && windPlayer !== null && shakeDetector !== null,
    load,
    play: () => player?.play(),
    stop: () => player?.stop(),
    energy,
    maxEnergy,
  }
}
