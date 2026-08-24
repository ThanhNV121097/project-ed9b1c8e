"use client";

import { useEffect, useState } from "react";
import styles from "./CenteredHelloWordPage.module.css";

type GreetingResponse = {
  greeting: {
    text: string;
  };
};

const apiBase = process.env.NEXT_PUBLIC_API_URL ?? "/api";

export default function CenteredHelloWordPage() {
  const [text, setText] = useState("");

  useEffect(() => {
    const controller = new AbortController();

    async function loadGreeting() {
      const response = await fetch(`${apiBase}/v1/greeting`, { signal: controller.signal });
      if (!response.ok) {
        return;
      }
      const data = (await response.json()) as GreetingResponse;
      setText(data.greeting.text);
    }

    void loadGreeting();
    return () => controller.abort();
  }, []);

  return (
    <main className={styles.page} aria-label={text}>
      <p className={styles.message}>{text}</p>
    </main>
  );
}
