export type GreetingResponse = {
  greeting: {
    text: string;
    updatedAt: string;
  } | null;
  error: {
    code: 'not_found' | 'internal_error';
    message: string;
  } | null;
};

export const greetingResponse: GreetingResponse = {
  greeting: {
    text: 'Hello Word',
    updatedAt: '2025-02-14T00:00:00Z',
  },
  error: null,
};
