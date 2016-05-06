var StemcellItem = React.createClass({
  render: function() {
    return (
      <div className="item">
        <div className="content">
          <a className="header">{this.props.name}</a>
          <div className="meta">
            <i className="pagelines icon"></i>
            <span className="category">{"v" + this.props.version}</span>
          </div>
        </div>
      </div>
    );
  }
});

var ReleaseItem = React.createClass({
  render: function() {
    return (
      <div className="item">
        <i className="cubes icon"></i>
        <div className="content">
          <a className="header">{this.props.name}</a>
          <div className="meta"><span className="category">{"v" + this.props.version}</span></div>
        </div>
      </div>
    );
  }
});

var DeploymentItem = React.createClass({
  render: function() {
    return (
        <div className="ui card">
          <div className="image">
            <img src="assets/images/bosh_deployment_logo.png"/>
          </div>

          <div className="content">
            <h1 className="ui header">{this.props.deployment.name}</h1>
          </div>

          <div className="content">
            <h2 className="ui header">Releases</h2>
            <div className="ui relaxed list">
            {this.props.deployment.releases.map(function(release) {
              return <ReleaseItem name={release.name} version={release.version}/>
            })}
            </div>
          </div>

          <div className="extra content">
            <h2 className="ui header">Stemcells</h2>
            <div classNameName="ui relaxed list">
            {this.props.deployment.stemcells.map(function(stemcell) {
              return <StemcellItem name={stemcell.name} version={stemcell.version}/>
            })}
            </div>
          </div>

          <div className="extra content">
            <span className="right floated time">
              <i className="history icon"></i>
              deployed {this.props.deployment.deploy_date}
            </span>
          </div>
         </div>
    );
  }
});

var DeploymentList = React.createClass({
  getInitialState: function() {
    return {deployments: []};
  },
  componentDidMount: function() {
    this.request = $.getJSON(this.props.url, function (result) {
      this.setState({deployments: result});
    }.bind(this));
  },
  componentWillUnmount: function() {
    this.request.abort();
  },
  render: function() {
    return (
     <div className="ui four stackable cards">
     {this.state.deployments.map(function(deployment) {
       deployment.deploy_date = jQuery.timeago(deployment.deploy_date);
       return <DeploymentItem deployment={deployment}/>
     })}
     </div>
    );
  }
});

var deployments = <DeploymentList url="/api/v1/deployments" />;
ReactDOM.render(deployments, document.getElementById('content'));
