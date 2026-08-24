import { centeredHelloWordPageMock } from "../lib/mock/centered-hello-word-page";
import styles from "./CenteredHelloWordPage.module.css";

export function CenteredHelloWordPage() {
  return (
    <main className={styles.page} aria-label={centeredHelloWordPageMock.greeting.text}>
      <p className={styles.message}>{centeredHelloWordPageMock.greeting.text}</p>
    </main>
  );
}
