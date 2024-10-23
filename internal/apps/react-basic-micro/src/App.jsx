import "./App.css";
import TodoList from "./components/TodoList.jsx";
import AddTodoForm from "./components/AddTodoForm.jsx";
import TodoSummary from "./components/TodoSummary.jsx";
import useItems from "./hooks/useTodos";

export default function App() {
  const {
    items,
    setItemCompleted,
    addItem,
    deleteItem,
    deleteAllCompletedItems,
  } = useItems();

  return (
    <>
      <main className="py-10 h-screen space-y-5 overflow-y-auto">
        <h1 className="font-bold text-3xl text-center">Your Todos</h1>
        <div className="max-w-lg mx-auto bg-slate-100 rounded-md p-5 space-y-6">
          <AddTodoForm onSubmit={addItem} />
          <TodoList
            items={items}
            onCompletedChange={setItemCompleted}
            onDelete={deleteItem}
          />
        </div>
        <TodoSummary
          items={items}
          deleteAllCompleted={deleteAllCompletedItems}
        />
      </main>
    </>
  );
}
