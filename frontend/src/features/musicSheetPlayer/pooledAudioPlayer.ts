const loadAudio = function (src: string): Promise<HTMLAudioElement> {
	return new Promise((resolve, reject) => {
		const audio = new Audio(src);
		audio.preload = "auto";
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
	private index = 0;
	soundFile: string;
	poolSize: number;

	constructor(soundFile: string, poolSize: number) {
		super();

		this.soundFile = soundFile;
		this.poolSize = poolSize;
	}

	async load() {
		this.audio = await Promise.all(
			Array.from({ length: this.poolSize }).map(() => loadAudio(this.soundFile))
		);
		this.dispatchEvent(new Event("load"));
	}

	play() {
		this.index = (this.index + 1) % this.audio.length;
		const audio = this.audio[this.index];

		audio.currentTime = 0;
		audio.play();

		console.log(
			"play",
			audio.currentSrc,
			this.index,
			audio.paused,
			audio.ended
		);
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
		player.load();
	});
};
