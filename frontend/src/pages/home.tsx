import { Container, Input, Stack, Text } from "@chakra-ui/react";
import { useTodoListQuery } from "@/hooks";

const Home = () => {
  const { data } = useTodoListQuery();
  const { data: todoList = [] } = data ?? {};

  return (
    <Container>
      <Stack>
        <Input placeholder="New Todo" />
        {todoList.map((todo) => (
          <Text key={todo.id}>{todo.title}</Text>
        ))}
      </Stack>
    </Container>
  );
};

export default Home;
