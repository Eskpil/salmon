import { useQuery } from "@tanstack/react-query";
import { Link, useNavigate } from "react-router-dom";
import { Layout } from "../../components/Layout";
import useTitle from "../../hooks/useTitle";
import { Node } from "../../types/node";

interface Props {}
interface NodeElementProps {
  node: Node;
}

const NodeElement: React.FC<NodeElementProps> = ({ node }) => {
  const navigate = useNavigate();

  return (
    <tr
      className="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50"
      onClick={() => navigate(`/nodes/${node.id}/overview`)}
    >
      <th
        scope="row"
        className="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white"
      >
        {node.hostname}
      </th>
      <td className="py-4 px-6">
        <span className="bg-green-100 text-green-800 text-xs font-semibold mr-2 px-2.5 py-0.5 rounded dark:bg-green-200 dark:text-green-900">
          Healthy
        </span>
      </td>
    </tr>
  );
};

export const NodesPage: React.FC<Props> = () => {
  const { data: nodes, isLoading } = useQuery<Node[]>(["nodes"], () =>
    fetch("http://cyan.local:8090/api/nodes/").then((res) => res.json())
  );

  useTitle("Nodes");

  if (isLoading) return <div>loading...</div>;

  return (
    <Layout>
      <div className="overflow-x-auto relative shadow-md sm:rounded-lg border col-span-6 col-start-2">
        <table className="w-full text-sm text-left text-gray-500 dark:text-gray-400">
          <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
            <tr>
              <th scope="col" className="py-3 px-6">
                Hostname
              </th>
              <th scope="col" className="py-3 px-6">
                Status
              </th>
            </tr>
          </thead>

          <tbody>
            {nodes?.map((node) => (
              <NodeElement node={node} />
            ))}
          </tbody>
        </table>
      </div>
    </Layout>
  );
};
