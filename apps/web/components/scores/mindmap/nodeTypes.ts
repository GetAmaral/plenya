import { GroupNode } from './GroupNode'
import { SubgroupNode } from './SubgroupNode'
import { ItemNode } from './ItemNode'
import { LevelNode } from './LevelNode'

export const nodeTypes = {
  groupNode: GroupNode,
  subgroupNode: SubgroupNode,
  itemNode: ItemNode,
  levelNode: LevelNode,
} as const
