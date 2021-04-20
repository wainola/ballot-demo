export const getInfoVoter = async (
  deployedInstance,
  accountSelected,
  chairperson,
  dispatcher
) => {
  try {
    const response = await deployedInstance.getInfoVoter(accountSelected, {
      from: chairperson,
    });
    console.log("GETTING INFO VOTER", response);
    dispatcher({
      type: "SET_LAST_INFO_VOTER",
      payload: response,
    });
  } catch (error) {
    console.log("Error on getting last voter", error);
  }
};
