export class PeriodicAudioPlayer extends EventTarget {
	audio: HTMLAudioElement | null = null;
	volume: number;
	startTime: number;
	src: string;
	private playing = false;

	constructor(src: string, startTime = 0, volume = 1) {
		super();
		this.volume = volume;
		this.startTime = startTime;
		this.src = src;
	}

	load() {
		this.audio = new Audio(this.src);
		this.audio.volume = this.volume;

		this.audio.addEventListener("canplaythrough", () => {
			this.dispatchEvent(new Event("load"));
		});
		this.audio.addEventListener("error", () => {
			this.dispatchEvent(new Event("error"));
		});
    this.dispatchEvent(new Event("load"));
	}

	play(interval: number, times: number) {
		if (!this.audio) return;
		if (this.playing) return;

		this.playing = true;
		let count = 0;

		const intervalId = setInterval(() => {
			count += 1;
			if (count > times) {
				clearInterval(intervalId);
				this.playing = false;
				return;
			}

			if (!this.audio) return;
			this.audio.currentTime = this.startTime;
			this.audio.play();
		}, interval);
	}
}

export const createPeriodicAudioPlayer = function (
	src: string,
	startTime = 0,
	volume = 1
): Promise<PeriodicAudioPlayer> {
	return new Promise((resolve, reject) => {
    const player = new PeriodicAudioPlayer(src, startTime, volume);
		player.addEventListener("load", () => {
			resolve(player);
		});
		player.addEventListener("error", () => {
			reject(player);
		});
		player.load();
    // resolve(player);
	});
};
