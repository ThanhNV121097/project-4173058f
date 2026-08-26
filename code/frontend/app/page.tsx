"use client";

import { greetingResponse } from '../lib/mock/display-stored-greeting';
import styles from '../components/GreetingFrame.module.css';

function GreetingFrame({
  greeting,
  note,
  label,
}: {
  greeting: string;
  note?: string;
  label: string;
}) {
  return (
    <section className={styles.shell}>
      <div className={styles.frame} aria-label={label}>
        <div className={styles.label}>{label}</div>
        <h1 className={styles.text}>{greeting}</h1>
        {note ? <p className={styles.note}>{note}</p> : null}
      </div>
    </section>
  );
}

export default function Home() {
  const greeting = greetingResponse.greeting?.text ?? 'Hello Word';

  return (
    <main>
      <GreetingFrame greeting={greeting} label="Loaded" />
      <GreetingFrame greeting="Hello Word" label="Loading" note="Loading stored greeting…" />
      <GreetingFrame greeting="Hello Word" label="Error" note="Could not read greeting row." />
    </main>
  );
}


