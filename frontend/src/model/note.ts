import {
	Note as ResponseNote,
	Pitch as ResponsePitch,
} from "~/sdk/api/music_sheet_pb";

export type Pitch = "C3" | "D3" | "E3" | "F3" | "G3" | "A4" | "B4" | "C4";

export type Note = {
	pitch: Pitch[];
};

export const PitchFromResponse = (pitch: ResponsePitch): Pitch => {
	switch (pitch) {
		case ResponsePitch.C3:
			return "C3";
		case ResponsePitch.D3:
			return "D3";
		case ResponsePitch.E3:
			return "E3";
		case ResponsePitch.F3:
			return "F3";
		case ResponsePitch.G3:
			return "G3";
		case ResponsePitch.A4:
			return "A4";
		case ResponsePitch.B4:
			return "B4";
		case ResponsePitch.C4:
			return "C4";
		case ResponsePitch.UNSPECIFIED:
			throw new Error("Unspecified pitch");
		default:
			throw new Error(pitch satisfies never);
	}
};

export const PitchToResponse = (pitch: Pitch): ResponsePitch => {
	switch (pitch) {
		case "C3":
			return ResponsePitch.C3;
		case "D3":
			return ResponsePitch.D3;
		case "E3":
			return ResponsePitch.E3;
		case "F3":
			return ResponsePitch.F3;
		case "G3":
			return ResponsePitch.G3;
		case "A4":
			return ResponsePitch.A4;
		case "B4":
			return ResponsePitch.B4;
		case "C4":
			return ResponsePitch.C4;
		default:
			throw new Error(pitch satisfies never);
	}
};

export const NoteFromResponse = (note: ResponseNote): Note => {
	return { pitch: note.pitches.map(PitchFromResponse) };
};

export const NoteToResponse = (note: Note): ResponseNote => {
	return new ResponseNote({
		pitches: note.pitch.map(PitchToResponse),
	});
};
