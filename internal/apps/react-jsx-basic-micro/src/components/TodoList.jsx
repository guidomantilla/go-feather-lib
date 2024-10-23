import TodoItem from "./TodoItem";
import PropTypes from "prop-types";

TodoList.propTypes = {
  items: PropTypes.array.isRequired,
  onCompletedChange: PropTypes.func.isRequired,
  onDelete: PropTypes.func.isRequired,
};

export default function TodoList({ items, onCompletedChange, onDelete }) {
  const itemsSorted = items.sort((a, b) => {
    if (a.completed === b.completed) {
      return b.id - a.id;
    }
    return a.completed ? 1 : -1;
  });

  return (
    <>
      <div className="space-y-2">
        {itemsSorted.map((item) => (
          <TodoItem
            key={item.id}
            item={item}
            onCompletedChange={onCompletedChange}
            onDelete={onDelete}
          />
        ))}
      </div>
      {items.length === 0 && (
        <p className="text-center text-sm text-gray-500">
          No todos yet. Add a new one above.
        </p>
      )}
    </>
  );
}
