import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { getJson, getTargets } from "@/index";
import { useState } from "react";

type Props = {
  count: number;
};

export const Counter = () => {
  const { count: initialCount } = getJson<Props>("counter");
  const [count, setCount] = useState(initialCount);
  return (
    <div className="grid gap-5">
      <p className="text-2xl font-bold">{count}</p>
      <div className="grid grid-cols-2 gap-5">
        <button
          className="border rounded-md px-2 py-1 bg-black text-white"
          onClick={() => setCount(count + 1)}
        >
          Increment
        </button>
        <button
          className="border rounded-md px-2 py-1 bg-black text-white"
          onClick={() => setCount(count - 1)}
        >
          Decrement
        </button>
      </div>
    </div>
  );
};

getTargets("counter").forEach((target) =>
  createRoot(target).render(
    <StrictMode>
      <Counter />
    </StrictMode>
  )
);
