import { FC } from "react";
import { Performer_findPerformer as Performer } from "src/graphql/definitions/Performer";

interface PerformerNameProps {
  performer: Pick<Performer, "name" | "disambiguation" | "deleted">;
  as?: string | null;
}

const PerformerName: FC<PerformerNameProps> = ({ performer, as }) => {
  if (!as)
    return (
      <>
        {performer.deleted ? (
          <del>{performer.name}</del>
        ) : (
          <span>{performer.name}</span>
        )}
        {performer.disambiguation && (
          <small className="ms-1 text-small text-muted">
            ({performer.disambiguation})
          </small>
        )}
      </>
    );
  return (
    <>
      <span>{as}</span>
      <small className="ms-1 text-small text-muted">
        ({performer.name})
        {performer.disambiguation && (
          <small className="ms-1 text-small text-muted">
            ({performer.disambiguation})
          </small>
        )}
      </small>
    </>
  );
};

export default PerformerName;
