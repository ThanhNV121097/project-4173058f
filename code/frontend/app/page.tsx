"use client";

import type { GreetingResponse } from '../lib/mock/display-stored-greeting';
import { greetingResponse } from '../lib/mock/display-stored-greeting';
import styles from '../components/GreetingFrame.module.css';

type GreetingState = 'loaded' | 'loading' | 'error';

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
    <section className={styles.shell} aria-label={label}>
      <div className={styles.frame}>
        <div className={styles.label}>{label}</div>
        <h1 className={styles.text}>{greeting}</h1>
        {note ? <p className={styles.note}>{note}</p> : null}
      </div>
    </section>
  );
}

function selectState(response: GreetingResponse): GreetingState {
  if (response.error) return 'error';
  if (!response.greeting?.text) return 'loading';
  return 'loaded';
}

export default function Home() {
  const view = selectState(greetingResponse);
  const greeting = greetingResponse.greeting?.text ?? 'Hello Word';

  if (view === 'loading') {
    return <GreetingFrame greeting="Hello Word" label="Loading" note="Loading stored greeting…" />;
  }

  if (view === 'error') {
    return <GreetingFrame greeting="Hello Word" label="Error" note="Could not read greeting row." />;
  }

  return <GreetingFrame greeting={greeting} label="Loaded" />;
}
