package constants

const ReactPromptTemplate = "请注意，你是一个有能力调用外部工具的智能助手。\n\n可用工具如下:\n%s\n\n请严格按照以下格式进行回应:\n\nThought: 你的思考过程，用于分析问题、拆解任务和规划下一步行动。\nAction: 你决定采取的行动，必须是以下格式之一:\n- %s[%s]`:调用一个可用工具。\n- Finish[最终答案]:当你认为已经获得最终答案时。\n- 当你收集到足够的信息，能够回答用户的最终问题时，你必须在Action:字段后使用 Finish[最终答案] 来输出最终答案。\n- 问你三次必须给出Finish[最终答案]\n\n现在，请开始解决以下问题:\nQuestion: %s\nHistory: %s"
const ToolsSearch = "Search"
const DouBaoFinish = "Finish"
const ThoughtMatch = "(?s)Thought:\\s*(.*?)\\s*Action:"
const ActionMatch = "Action:\\s*(.+)"
const ToolMatch = "^(\\w+)\\[(.*)\\]$"
const FinishMatch = "Finish\\s*(.+)"
