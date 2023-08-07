export type Paginated<T> = {
  cursor: string;
  data: T[];
};
