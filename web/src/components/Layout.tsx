import { useNavigate } from "react-router-dom";
import { Navbar } from "./Navbar";

interface Props {
  children: React.ReactNode;
  customButton?: React.ReactNode | null;
}

export const Layout: React.FC<Props> = ({ children, customButton }) => {
  const navigate = useNavigate();
  return (
    <div>
      <Navbar />
      <div className="grid grid-cols-8 m-8">
        <div className="col-start-2 col-span-6">
          <div className="flex flex-row justify-between">
            <div>
              <button
                type="button"
                className="py-2.5 px-5 mr-2 mb-2 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-200 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700"
                onClick={() => navigate(-1)}
              >
                Back
              </button>
              <button
                type="button"
                className="py-2.5 px-5 mr-2 mb-2 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-200 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700"
                onClick={() => navigate(1)}
              >
                Forward
              </button>
            </div>
            {customButton != null ? <div>{customButton}</div> : null}
          </div>
          {children}
        </div>
      </div>
    </div>
  );
};
