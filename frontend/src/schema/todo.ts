import { Paginated } from "@/schema/common";

export type Todo = {
  id: string;
  title: string;
  isCompleted: boolean;
  completedAt: string | null;
  createdAt: string;
};

export type PaginatedTodo = Paginated<Todo>;
