import { useMutation } from "@tanstack/react-query";
import { queryClient } from "~/queryClient";
import {
	createV1,
	getV1,
} from "~/sdk/api/music_sheet-MusicSheetService_connectquery";

export const usePostMusicSheet = () => {
	const { mutationFn } = createV1.useMutation();
	return { postMusicSheet: mutationFn };
};
