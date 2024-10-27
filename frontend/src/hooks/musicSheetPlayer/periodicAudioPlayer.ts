export class PeriodicAudioPlayer extends EventTarget {
	audio: HTMLAudioElement;
	startTime: number;

	constructor(src: string, startTime = 0) {
		super();

		this.audio = new Audio(src);
		this.startTime = startTime;

		this.audio.addEventListener("canplaythrough", () => {
			this.dispatchEvent(new Event("load"));
		});
		this.audio.addEventListener("error", () => {
			this.dispatchEvent(new Event("error"));
		});
	}

	play(interval: number, times: number) {
		let count = 0;

		this.audio.pause();

		const intervalId = setInterval(() => {
			count += 1;
			if (count > times) {
				clearInterval(intervalId);
				return;
			}

			this.audio.currentTime = this.startTime;
			this.audio.play();
		}, interval);
	}
}

export const createPeriodicAudioPlayer = function (
	src: string,
	startTime = 0
): Promise<PeriodicAudioPlayer> {
	return new Promise((resolve, reject) => {
		const player = new PeriodicAudioPlayer(src, startTime);
		player.addEventListener("load", () => {
			resolve(player);
		});
		player.addEventListener("error", () => {
			reject(player);
		});
	});
};
