import { useEffect, useState } from "react";
import { dummyItems } from "../data/todos";

export default function useItems() {
  const [items, setItems] = useState(() => {
    const savedItems = JSON.parse(localStorage.getItem("items") || "[]");
    return savedItems.length > 0 ? savedItems : dummyItems;
  });

  useEffect(() => {
    localStorage.setItem("items", JSON.stringify(items));
  }, [items]);

  function setItemCompleted(id, completed) {
    setItems((prevTodos) =>
      prevTodos.map((item) => (item.id === id ? { ...item, completed } : item)),
    );
  }

  function addItem(title) {
    setItems((prevItems) => [
      {
        id: Date.now(),
        title,
        completed: false,
      },
      ...prevItems,
    ]);
  }

  function deleteItem(id) {
    setItems((prevItems) => prevItems.filter((item) => item.id !== id));
  }

  function deleteAllCompletedItems() {
    setItems((prevItems) => prevItems.filter((item) => !item.completed));
  }

  return {
    items,
    setItemCompleted,
    addItem,
    deleteItem,
    deleteAllCompletedItems,
  };
}
