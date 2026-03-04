package constants

const ReactPromptTemplate = "请注意，你是一个有能力调用外部工具的智能助手。\n\n可用工具如下:\n%s\n\n请严格按照以下格式进行回应:\n\nThought: 你的思考过程，用于分析问题、拆解任务和规划下一步行动。\nAction: 你决定采取的行动，必须是以下格式之一:\n- %s[%s]`:调用一个可用工具。\n- Finish[最终答案]:当你认为已经获得最终答案时。\n- 当你收集到足够的信息，能够回答用户的最终问题时，你必须在Action:字段后使用 Finish[最终答案] 来输出最终答案。\n- 问你三次必须给出Finish[最终答案]\n\n现在，请开始解决以下问题:\nQuestion: %s\nHistory: %s"
const ToolsSearch = "Search"
const DouBaoFinish = "Finish"
const ThoughtMatch = "(?s)Thought:\\s*(.*?)\\s*Action:"
const ActionMatch = "Action:\\s*(.+)"
const ToolMatch = "^(\\w+)\\[(.*)\\]$"
const FinishMatch = "Finish\\[([\\s\\S]*?)\\]"
const PlannerMatch = "\"([^\"]+)\""
const PlannerPromptTemplate = "你是一个顶级的AI规划专家。你的任务是将用户提出的复杂问题分解成一个由多个简单步骤组成的行动计划。\n请确保计划中的每个步骤都是一个独立的、可执行的子任务，并且严格按照逻辑顺序排列。\n你的输出必须是一个Python列表，其中每个元素都是一个描述子任务的字符串。\n\n问题: %s\n\n请严格按照以下格式输出你的计划,```go与```作为前后缀是必要的:\n```go\n[\"步骤1\", \"步骤2\", \"步骤3\", ...]\n```"
const ExecutorPromptTemplate = "你是一位顶级的AI执行专家。你的任务是严格按照给定的计划，一步步地解决问题。\n你将收到原始问题、完整的计划、以及到目前为止已经完成的步骤和结果。\n请你专注于解决“当前步骤”，并仅输出该步骤的最终答案，不要输出任何额外的解释或对话。\n\n# 原始问题:\n%s\n\n# 完整计划:\n%s\n\n# 历史步骤与结果:\n%s\n\n# 当前步骤:\n%d\n\n请仅输出针对“当前步骤”的回答:"
