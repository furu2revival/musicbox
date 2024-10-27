import { QueryClient } from "@tanstack/react-query";

export const queryClient = new QueryClient({
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
