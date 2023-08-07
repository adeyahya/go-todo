import { Paginated } from "@/schema/common";

export type Todo = {
  id: string;
  title: string;
  isCompleted: boolean;
  createdAt: string;
};

export type PaginatedTodo = Paginated<Todo>;
