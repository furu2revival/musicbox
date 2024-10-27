import type { MusicSheet as ResponseMusicSheet } from "~/sdk/api/music_sheet_pb";
import { type Note, NoteFromResponse } from "./note";

export type MusicSheet = {
	id: string;
	title: string;
	notes: Note[];
};

export const MusicSheetFromResponse = (
	musicSheet: ResponseMusicSheet
): MusicSheet => {
	return {
		id: musicSheet.musicSheetId,
		title: musicSheet.title,
		notes: musicSheet.notes.map(NoteFromResponse),
	};
};
