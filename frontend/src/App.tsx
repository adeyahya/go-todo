import { ChakraProvider } from "@chakra-ui/react";
import { RouterProvider } from "react-router-dom";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import routes from "@/routes";

const queryClient = new QueryClient();

function App() {
  return (
    <ChakraProvider>
      <QueryClientProvider client={queryClient}>
        <RouterProvider router={routes} />
      </QueryClientProvider>
    </ChakraProvider>
  );
}

export default App;
