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
  greeting: null,
  error: {
    code: 'not_found',
    message: 'Could not read greeting.',
  },
};
