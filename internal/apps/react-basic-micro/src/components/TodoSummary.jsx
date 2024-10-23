import PropTypes from "prop-types";

TodoSummary.propTypes = {
  items: PropTypes.array.isRequired,
  deleteAllCompleted: PropTypes.func.isRequired,
};

export default function TodoSummary({ items, deleteAllCompleted }) {
  const completedItems = items.filter((item) => item.completed);

  return (
    <div className="text-center space-y-2">
      <p className="text-sm font-medium">
        {completedItems.length}/{items.length} items completed
      </p>
      {completedItems.length > 0 && (
        <button
          onClick={deleteAllCompleted}
          className="text-red-500 hover:underline text-sm font-medium"
        >
          Delete all completed
        </button>
      )}
    </div>
  );
}
