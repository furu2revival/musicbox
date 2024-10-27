import { TransportProvider } from "@bufbuild/connect-query";
import { createConnectTransport } from "@connectrpc/connect-web";
import { QueryClientProvider } from "@tanstack/react-query";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import App from "./App.tsx";
import { queryClient } from "./queryClient.ts";
import "./main.module.css";

const finalTransport = createConnectTransport({
	baseUrl: "https://musicbox.averak.net",
	// 開発期間は、binary ではなく JSON 形式を用いる。
	// 理由: ブラウザの開発者ツールで human-readable なリクエスト/レスポンスメッセージを確認できるため。
	useBinaryFormat: false,
});

// biome-ignore lint/style/noNonNullAssertion: <explanation>
createRoot(document.getElementById("root")!).render(
	<StrictMode>
		<TransportProvider transport={finalTransport}>
			<QueryClientProvider client={queryClient}>
				<App />
			</QueryClientProvider>
		</TransportProvider>
	</StrictMode>,
);
