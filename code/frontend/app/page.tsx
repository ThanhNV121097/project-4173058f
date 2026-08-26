"use client";

import type { GreetingResponse } from '../lib/mock/display-stored-greeting';
import { greetingResponse } from '../lib/mock/display-stored-greeting';

type GreetingState = 'loaded' | 'loading' | 'error';

const state: GreetingState = 'loaded';

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
    <section className="greeting-shell" aria-label={label}>
      <div className="greeting-frame">
        <div className="greeting-label">{label}</div>
        <h1 className="greeting-text">{greeting}</h1>
        {note ? <p className="greeting-note">{note}</p> : null}
      </div>
    </section>
  );
}

function selectState(response: GreetingResponse): GreetingState {
  if (response.error) return 'error';
  if (!response.greeting?.text) return 'loading';
  return state;
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
