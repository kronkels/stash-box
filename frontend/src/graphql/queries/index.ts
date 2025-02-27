import { useContext } from "react";
import {
  useQuery,
  useLazyQuery,
  QueryHookOptions,
  LazyQueryHookOptions,
} from "@apollo/client";

import AuthContext from "src/AuthContext";
import { isAdmin } from "src/utils";

import { Category, CategoryVariables } from "../definitions/Category";
import { Categories } from "../definitions/Categories";
import { Edit, EditVariables } from "../definitions/Edit";
import { Edits, EditsVariables } from "../definitions/Edits";
import { Me } from "../definitions/Me";
import { Performer, PerformerVariables } from "../definitions/Performer";
import {
  FullPerformer,
  FullPerformerVariables,
} from "../definitions/FullPerformer";
import { Performers, PerformersVariables } from "../definitions/Performers";
import { Scene, SceneVariables } from "../definitions/Scene";
import { Scenes, ScenesVariables } from "../definitions/Scenes";
import {
  ScenesWithoutCount,
  ScenesWithoutCountVariables,
} from "../definitions/ScenesWithoutCount";
import { SearchAll, SearchAllVariables } from "../definitions/SearchAll";
import {
  SearchPerformers,
  SearchPerformersVariables,
} from "../definitions/SearchPerformers";
import { Studio, StudioVariables } from "../definitions/Studio";
import { Studios, StudiosVariables } from "../definitions/Studios";
import { Tag, TagVariables } from "../definitions/Tag";
import { Tags, TagsVariables } from "../definitions/Tags";
import { User, UserVariables } from "../definitions/User";
import { Users, UsersVariables } from "../definitions/Users";
import { PublicUser, PublicUserVariables } from "../definitions/PublicUser";
import { Config } from "../definitions/Config";
import { Version } from "../definitions/Version";
import {
  PendingEditsCount,
  PendingEditsCountVariables,
} from "../definitions/PendingEditsCount";
import { Site, SiteVariables } from "../definitions/Site";
import { Sites } from "../definitions/Sites";
import { Draft, DraftVariables } from "../definitions/Draft";
import { Drafts } from "../definitions/Drafts";

import CategoryQuery from "./Category.gql";
import CategoriesQuery from "./Categories.gql";
import EditQuery from "./Edit.gql";
import EditsQuery from "./Edits.gql";
import MeQuery from "./Me.gql";
import PerformerQuery from "./Performer.gql";
import FullPerformerQuery from "./FullPerformer.gql";
import PerformersQuery from "./Performers.gql";
import SceneQuery from "./Scene.gql";
import ScenesQuery from "./Scenes.gql";
import ScenesWithoutCountQuery from "./ScenesWithoutCount.gql";
import SearchAllQuery from "./SearchAll.gql";
import SearchPerformersQuery from "./SearchPerformers.gql";
import StudioQuery from "./Studio.gql";
import StudiosQuery from "./Studios.gql";
import TagQuery from "./Tag.gql";
import TagsQuery from "./Tags.gql";
import UserQuery from "./User.gql";
import UsersQuery from "./Users.gql";
import PublicUserQuery from "./PublicUser.gql";
import ConfigQuery from "./Config.gql";
import VersionQuery from "./Version.gql";
import PendingEditsCountQuery from "./PendingEditsCount.gql";
import SiteQuery from "./Site.gql";
import SitesQuery from "./Sites.gql";
import DraftQuery from "./Draft.gql";
import DraftsQuery from "./Drafts.gql";

export const useCategory = (variables: CategoryVariables, skip = false) =>
  useQuery<Category, CategoryVariables>(CategoryQuery, {
    variables,
    skip,
  });

export const useCategories = () => useQuery<Categories>(CategoriesQuery);

export const useEdit = (variables: EditVariables) =>
  useQuery<Edit, EditVariables>(EditQuery, {
    variables,
  });

export const useEdits = (variables: EditsVariables) =>
  useQuery<Edits, EditsVariables>(EditsQuery, {
    variables,
  });

export const useMe = (options?: QueryHookOptions<Me>) =>
  useQuery<Me>(MeQuery, options);

export const usePerformer = (variables: PerformerVariables, skip = false) =>
  useQuery<Performer, PerformerVariables>(PerformerQuery, {
    variables,
    skip,
  });

export const useFullPerformer = (variables: PerformerVariables, skip = false) =>
  useQuery<FullPerformer, FullPerformerVariables>(FullPerformerQuery, {
    variables,
    skip,
  });

export const usePerformers = (variables: PerformersVariables) =>
  useQuery<Performers, PerformersVariables>(PerformersQuery, {
    variables,
  });

export const useScene = (variables: SceneVariables, skip = false) =>
  useQuery<Scene, SceneVariables>(SceneQuery, {
    variables,
    skip,
  });

export const useScenes = (variables: ScenesVariables, skip = false) =>
  useQuery<Scenes, ScenesVariables>(ScenesQuery, {
    variables,
    skip,
  });

export const useScenesWithoutCount = (
  variables: ScenesVariables,
  skip = false
) =>
  useQuery<ScenesWithoutCount, ScenesWithoutCountVariables>(
    ScenesWithoutCountQuery,
    {
      variables,
      skip,
    }
  );

export const useSearchAll = (variables: SearchAllVariables, skip = false) =>
  useQuery<SearchAll, SearchAllVariables>(SearchAllQuery, {
    variables,
    skip,
  });

export const useSearchPerformers = (variables: SearchPerformersVariables) =>
  useQuery<SearchPerformers, SearchPerformersVariables>(SearchPerformersQuery, {
    variables,
  });

export const useLazySearchAll = (
  options?: LazyQueryHookOptions<SearchAll, SearchAllVariables>
) => useLazyQuery(SearchAllQuery, options);

export const useLazySearchPerformers = (
  options?: LazyQueryHookOptions<SearchPerformers, SearchPerformersVariables>
) => useLazyQuery(SearchPerformersQuery, options);

export const useStudio = (variables: StudioVariables, skip = false) =>
  useQuery<Studio, StudioVariables>(StudioQuery, {
    variables,
    skip,
  });

export const useStudios = (variables: StudiosVariables) =>
  useQuery<Studios, StudiosVariables>(StudiosQuery, {
    variables,
  });

export const useLazyStudios = (
  options?: LazyQueryHookOptions<Studios, StudiosVariables>
) => useLazyQuery(StudiosQuery, options);

export const useTag = (variables: TagVariables, skip = false) =>
  useQuery<Tag, TagVariables>(TagQuery, {
    variables,
    skip,
  });

export const useTags = (variables: TagsVariables) =>
  useQuery<Tags, TagsVariables>(TagsQuery, {
    variables,
  });
export const useLazyTags = (
  options?: LazyQueryHookOptions<Tags, TagsVariables>
) => useLazyQuery(TagsQuery, options);

export const usePrivateUser = (variables: UserVariables, skip = false) =>
  useQuery<User, UserVariables>(UserQuery, {
    variables,
    skip,
  });
export const usePublicUser = (variables: PublicUserVariables, skip = false) =>
  useQuery<PublicUser, PublicUserVariables>(PublicUserQuery, {
    variables,
    skip,
  });

export const useUser = (variables: UserVariables, skip = false) => {
  const Auth = useContext(AuthContext);
  const isUser = () => Auth.user?.name === variables.name;
  const showPrivate = isUser() || isAdmin(Auth.user);

  const privateUser = usePrivateUser(variables, skip || !showPrivate);
  const publicUser = usePublicUser(variables, skip || showPrivate);

  return showPrivate ? privateUser : publicUser;
};

export const useUsers = (variables: UsersVariables) =>
  useQuery<Users, UsersVariables>(UsersQuery, {
    variables,
  });

export const useConfig = () => useQuery<Config>(ConfigQuery);

export const useVersion = () => useQuery<Version>(VersionQuery);

export const usePendingEditsCount = (variables: PendingEditsCountVariables) =>
  useQuery<PendingEditsCount, PendingEditsCountVariables>(
    PendingEditsCountQuery,
    { variables }
  );

export const useSite = (variables: SiteVariables, skip = false) =>
  useQuery<Site, SiteVariables>(SiteQuery, {
    variables,
    skip,
  });

export const useSites = () => useQuery<Sites>(SitesQuery);

export const useDraft = (variables: DraftVariables, skip = false) =>
  useQuery<Draft, DraftVariables>(DraftQuery, {
    variables,
    skip,
  });

export const useDrafts = () => useQuery<Drafts>(DraftsQuery);
