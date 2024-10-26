import type { Note } from "./note";

export type MusicSheet = {
	id: string;
	title: string;
	notes: Note[];
};
