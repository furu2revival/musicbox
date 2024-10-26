// 0: unspecified
// 1: C3, 2: D3, ...
export type Pitch = 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8;

export type Note = {
	pitch: Pitch[];
};
