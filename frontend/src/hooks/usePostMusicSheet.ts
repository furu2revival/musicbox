import { useMutation } from "@tanstack/react-query";
import { queryClient } from "~/queryClient";
import {
	createV1,
	getV1,
} from "~/sdk/api/music_sheet-MusicSheetService_connectquery";

export const usePostMusicSheet = () => {
	const { data, isLoading, mutate } = useMutation(createV1.useMutation, {
		onSuccess: () => {
			queryClient.invalidateQueries(getV1.getQueryKey());
		},
	});
	return { postMusicSheet: mutate, resData: data, isLoading };
};
