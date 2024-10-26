const isPlaying = function (audio: HTMLAudioElement) {
	return audio.currentTime > 0 && !audio.paused && !audio.ended;
};

const loadAudio = function (src: string): Promise<HTMLAudioElement> {
	return new Promise((resolve, reject) => {
		const audio = new Audio(src);
		audio.addEventListener("canplaythrough", () => {
			resolve(audio);
		});
		audio.addEventListener("error", () => {
			reject("Failed to load audio file");
		});
	});
};

export class PooledAudioPlayer extends EventTarget {
	private audio: HTMLAudioElement[] = [];

	constructor(soundFile: string, poolSize: number) {
		super();
		this.load(soundFile, poolSize);
	}

	private async load(src: string, count: number) {
		this.audio = await Promise.all(
			Array.from({ length: count }).map(() => loadAudio(src))
		);
		this.dispatchEvent(new Event("load"));
	}

	play() {
		const audio = this.audio.find((audio) => !isPlaying(audio));
		if (!audio) return;

		audio.play();
	}
}

export const createPooledAudioPlayer = function (
	soundFile: string,
	poolSize: number
): Promise<PooledAudioPlayer> {
	return new Promise((resolve) => {
		const player = new PooledAudioPlayer(soundFile, poolSize);
		player.addEventListener("load", () => {
			resolve(player);
		});
	});
};
