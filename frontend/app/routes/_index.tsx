import type { MetaFunction } from "@remix-run/node";
import MyButton from "../components/button";

export const meta: MetaFunction = () => {
  return [
    { title: "New Remix App" },
    { name: "description", content: "Welcome to Remix!" },
  ];
};

export default function Index() {
  return (
    <div >
      <MyButton />
    </div>
  );
}
