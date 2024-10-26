import { TransportProvider } from "@bufbuild/connect-query";
import { createConnectTransport } from "@connectrpc/connect-web";
import type { MetaFunction } from "@remix-run/node";
import MyButton from "../components/button";
import { QueryClient } from "@tanstack/query-core";
import {QueryClientProvider} from "@tanstack/react-query";

export const meta: MetaFunction = () => {
	return [
		{ title: "New Remix App" },
		{ name: "description", content: "Welcome to Remix!" },
	];
};

const finalTransport = createConnectTransport({
	baseUrl: "http://43.207.202.40:8000",
	// 開発期間は、binary ではなく JSON 形式を用いる。
	// 理由: ブラウザの開発者ツールで human-readable なリクエスト/レスポンスメッセージを確認できるため。
	useBinaryFormat: false,
});

const queryClient = new QueryClient({
	defaultOptions: {
		queries: {
			queryKeyHashFn: (object: unknown) =>
				JSON.stringify(object, (_, value) => {
					// BigInt 型は JSON に変換できないので、文字列に変換する。
					return typeof value === "bigint" ? value.toString() : value;
				}),
		},
	},
});

export default function Index() {
	return (
		<TransportProvider transport={finalTransport}>
			<QueryClientProvider client={queryClient}>
				<MyButton />
			</QueryClientProvider>
		</TransportProvider>
	);
}
