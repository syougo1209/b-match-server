export const range = (start: number, end: number) => [...new Array(end - start).keys()].map((n) => n + start)
