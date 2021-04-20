pragma solidity 0.8.3;

contract FactoryBallot {
    address public chairperson;

    struct Proposal {
        string title;
        string description;
        uint256 votes;
        uint256 weight;
    }

    struct Voter {
        string name;
        string lastname;
        string email;
        uint256 weight;
        Proposal[] proposals;
    }

    mapping(address => Voter) public voters;

    constructor() {
        chairperson = msg.sender;
    }

    modifier onlyChairperson {
        require(
            msg.sender == chairperson,
            "Only chairperson can register new voters"
        );
        _;
    }

    modifier checkInfoVoter {
        require(
            msg.sender == chairperson,
            "Only chairperson can see infor voter"
        );
        _;
    }

    modifier checkAddressNotEmpty(address v) {
        require(bytes20(v).length != 0, "No address supplied");
        _;
    }

    modifier checkPay {
        require(msg.value != 0, "You must pay for this transaction");
        _;
    }

    modifier checkProposalTitleNotEmpty(string memory proposalTitle) {
        bytes memory p = bytes(proposalTitle);
        require(p.length != 0, "You must provide the proposal title");
        _;
    }

    modifier checkNotChairperson {
        require(
            msg.sender != chairperson,
            "Chairperson doesn't have proposals of his own"
        );
        _;
    }

    modifier checkAddressHasProposals(address addr) {
        Voter memory v = voters[addr];
        Proposal[] memory p = v.proposals;
        require(p.length != 0, "Account doesn't have any proposals");
        _;
    }

    // this method could be placed on a library file
    function convertStringToBytes32(bytes memory bytesToConver)
        internal
        pure
        returns (bytes32 dataConverted)
    {
        dataConverted = keccak256(bytesToConver);
        return dataConverted;
    }

    // CHAIRPERONS REGISTER VOTERS
    function registerVoters(
        address voter,
        string memory name,
        string memory lastname,
        string memory email
    ) public onlyChairperson checkAddressNotEmpty(voter) {
        Voter storage v = voters[voter];
        v.name = name;
        v.lastname = lastname;
        v.email = email;
    }

    function getInfoVoter(address voter)
        public
        view
        checkInfoVoter
        checkAddressNotEmpty(voter)
        returns (
            string memory name,
            string memory lastname,
            string memory email
        )
    {
        Voter storage v = voters[voter];
        name = v.name;
        lastname = v.lastname;
        email = v.email;

        return (name, lastname, email);
    }

    // set the proposal and it could contain weight in order
    // to increase the chances to be on top of the proposal list
    function setProposal(string memory title, string memory description)
        public
        payable
        checkNotChairperson
        checkPay
    {
        Voter storage voter = voters[msg.sender];

        // note: msg.value represent also the weigth of the proposal
        Proposal memory p = Proposal(title, description, 0, msg.value);
        voter.proposals.push(p);
    }

    function getProposal(address addr, string memory proposalTitle)
        public
        view
        checkNotChairperson
        checkAddressNotEmpty(addr)
        checkProposalTitleNotEmpty(proposalTitle)
        returns (Proposal memory result)
    {
        Voter memory v = voters[addr];
        for (uint256 i = 0; i < v.proposals.length; i++) {
            // we do this becauses is the only way
            // as far as I know, to compare to strings
            bytes32 str1 = keccak256(bytes(v.proposals[i].title));
            bytes32 str2 = keccak256(bytes(proposalTitle));
            if (str1 == str2) {
                result = v.proposals[i];
                return result;
            }
        }
    }

    function vote(address addr, string memory proposalTitle)
        public
        payable
        checkAddressNotEmpty(addr)
        checkAddressHasProposals(addr)
        checkPay
    {
        Voter storage voter = voters[addr];
        Proposal[] storage proposals = voter.proposals;

        bytes32 proposalTitleForComparison = keccak256(bytes(proposalTitle));
        for (uint256 i = 0; i < proposals.length; i++) {
            bytes32 proposalOfVoter = keccak256(bytes(proposals[i].title));
            if (proposalTitleForComparison == proposalOfVoter) {
                proposals[i].votes += 1;
                break;
            }
        }
    }

    function getVotesForProposal(address addr, string memory proposalTitle)
        public
        view
        returns (uint256 voteCount)
    {
        Voter memory voter = voters[addr];
        Proposal[] memory proposals = voter.proposals;

        bytes32 proposalTitleConverted =
            convertStringToBytes32(bytes(proposalTitle));
        for (uint256 i = 0; i < proposals.length; i++) {
            bytes32 proposalVoterTitleConverted =
                convertStringToBytes32(bytes(proposals[i].title));
            if (proposalVoterTitleConverted == proposalTitleConverted) {
                return proposals[i].votes;
            }
        }
    }

    function getOwnProposals()
        public
        view
        checkNotChairperson
        checkAddressNotEmpty(msg.sender)
        returns (Proposal[] memory ownProposals)
    {
        ownProposals = voters[msg.sender].proposals;
        if (ownProposals.length == 0) {
            revert("Not proposals for the address, operation reverted");
        }
        return ownProposals;
    }
}
