export class AudioContextPlayer extends EventTarget {
	src: string;
	private context: AudioContext | null = null;
	private buffer: AudioBuffer | null = null;
	private nextNode: AudioBufferSourceNode | null = null;

	constructor(src: string) {
		super();

		this.src = src;
	}

	async load() {
		this.context = new AudioContext();

		const res = await fetch(this.src);
		if (!res.ok) {
			throw new Error("Failed to load audio file");
		}
		const data = await res.arrayBuffer();
		this.buffer = await this.context.decodeAudioData(data);

		await this.prepareNextNode();
		this.dispatchEvent(new Event("load"));
	}

	async prepareNextNode() {
		if (!this.context) {
			throw new Error("Context is not loaded");
		}
		this.nextNode = this.context.createBufferSource();
		this.nextNode.buffer = this.buffer;
		this.nextNode.loop = false;
		this.nextNode.connect(this.context.destination);
	}

	async play() {
		if (!this.nextNode) {
			throw new Error("Node is not prepared");
		}

		this.nextNode.start();
		this.prepareNextNode();
	}
}

export const createAudioContextPlayer = function (
	src: string
): Promise<AudioContextPlayer> {
	return new Promise((resolve) => {
		const player = new AudioContextPlayer(src);
		player.addEventListener("load", () => {
			resolve(player);
		});
		player.load();
	});
};
