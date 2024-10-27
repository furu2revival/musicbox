import { MusicSheetEditor } from "./components/MusicSheetEditor";
import { useMusicSheet } from "./hooks/useMusicSheet";
import { MusicSheetFromResponse } from "./model/musicSheet";
import { NoteFromResponse } from "./model/note";

function App() {
	const currentUrl = new URL(window.location.href);
	const musicSheetId = currentUrl.searchParams.get("musicSheetId");
	const { data, isError, isLoading } = useMusicSheet(musicSheetId);
	// 存在しない musicSheetId が指定された場合、トップページにリダイレクトする
	if (isLoading) {
		return <div>loading...</div>;
	}
	if (isError) {
		window.location.href = "/";
	}

	return (
		<MusicSheetEditor
			maxNotes={64}
			musicSheet={
				data?.musicSheet ? MusicSheetFromResponse(data.musicSheet) : undefined
			}
			energy={20}
		/>
	);
}

export default App;
